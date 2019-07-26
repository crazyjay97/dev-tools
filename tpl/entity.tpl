package {{packageName}}.{{moduleName}}.entity;

import com.baomidou.mybatisplus.annotations.TableId;
import com.baomidou.mybatisplus.annotations.TableName;
import lombok.Data;


/**
 * {{ table.TableComment }}
 *
 * @author {{ authorName }}
 * @email {{ emailAddress }}
 * @date {{ genTime }}
 */
@TableName("{{ table.TableName }}")
@Data
public class {{ className }}Entity implements Serializable {
	private static final long serialVersionUID = 1L;
    {% for column in listColumns %}{% if !column.IsJoinColumn%}
	/**
	 * {{ column.ColumnComment }}
	 */{% if column.ColumnKey == "PRI" %}
    @TableId {% endif %}
    private {{ column.JavaType }} {{ column.FieldName}};{% endif %}{% endfor %}{% for column in listColumns %}{% if column.IsJoinColumn %}
	/**
	 * {{ column.ColumnComment }}
	 */
    @TableField(exist = false)
    private {{ column.JavaType }} {{ column.FieldName}};{% endif %}{% endfor %}
}
