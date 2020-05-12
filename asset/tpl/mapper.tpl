<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">

<mapper namespace="{{ packageName }}.{{ moduleName }}.dao.{{ className }}Domain">

    <resultMap type="{{ packageName }}.{{ moduleName }}.domain.{{ className }}Domain" id="baseMap">
        {% for column in listColumns %}
        <result property="{{ column.FieldName }}" column="{{ column.ColumnName }}"/>
        {% endfor %}
    </resultMap>


    <select id="page" parameterType="object" resultMap="baseMap">
        SELECT
        {% for column in columns %}
        t1.{{ column.ColumnName }}{% if !forloop.Last %},
        {% endif %}
        {% endfor %}
        {% for column in joinTables %}{% if forloop.First %},
        {% endif %}
        t{{ forloop.Counter+1 }}.{{ column.SearchColumn }} AS {{ column.Alias }}
        {% if !forloop.Last %},
        {% endif %}
        {% endfor %}
        FROM {{ table.TableName }} t1
        {% for joinTable in joinTables %}
        LEFT JOIN {{ joinTable.TableName }} t{{ forloop.Counter+1 }} ON t1.{{ joinTable.SelfColumn }} = t{{ forloop.Counter+1 }}.{{ joinTable.JoinColumn }}
        {% endfor %}
        WHERE {% if table.LogicDel %} t1.deleted = 0 {% else %} 1 = 1 {% endif %}

        {% for column in columns %}
        <if test="{{ column.FieldName }} != null and {{ column.FieldName }} !='' ">
            AND t1.{{ column.ColumnName }} = {{"#{"}}{{ column.FieldName }}{{"}"}}
        </if>
        {% endfor %}
    </select>

{% if table.LogicDel %}
    <update id="deleteBatchIds" parameterType="java.util.List">
        UPDATE {{ table.TableName }} SET deleted = 1
        WHERE ID IN
        <foreach collection="list" index="index" item="item" open="(" separator="," close=")">
            #{item}
        </foreach>
    </update>
{% endif %}
</mapper>