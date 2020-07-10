package {{packageName}}.{{moduleName}}.dao;

import {{packageName}}.{{moduleName}}.domain.{{ className }}Domain;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import org.apache.ibatis.annotations.Mapper;

import java.util.List;
import java.util.Map;

/**
  * {{ table.TableComment }}
  *
  * @author {{ authorName }}
  * @email {{ emailAddress }}
  * @date {{ genTime }}
  */
@Mapper
public interface {{ className }}Dao extends BaseMapper<{{ className }}Domain> {

{% if table.LogicDel %}
    void deleteBatchIds(List<String> ids);
{% endif %}

    List<{{ className }}Domain> page(Page<{{ className }}Domain> page,Map params);

}